package main

import (
	"fmt"

	"bazil.org/fuse"
)

func main() {
	c, err := fuse.Mount("/home/paschal/fs", fuse.FSName("p"))
	if err != nil {
		panic(err)
	}
	defer fuse.Unmount("/home/paschal/fs")
	defer c.Close()

	reqChan := make(chan fuse.Request)

	go func() {
		for req := range reqChan {
			fmt.Println("req is -> ", req, err)

			switch r := req.(type) {
			case *fuse.AccessRequest,
				*fuse.BatchForgetRequest, *fuse.CreateRequest, *fuse.DestroyRequest,
				*fuse.FAllocateRequest, *fuse.FlushRequest, *fuse.ForgetRequest,
				*fuse.FsyncRequest, *fuse.GetattrRequest, *fuse.GetxattrRequest,
				*fuse.InterruptRequest, *fuse.LinkRequest, *fuse.ListxattrRequest,
				*fuse.LockRequest, *fuse.LockWaitRequest, *fuse.LookupRequest,
				*fuse.MkdirRequest, *fuse.MknodRequest, *fuse.OpenRequest,
				*fuse.PollRequest, *fuse.QueryLockRequest, *fuse.ReadRequest,
				*fuse.ReadlinkRequest, *fuse.ReleaseRequest, *fuse.RemoveRequest,
				*fuse.RemovexattrRequest, *fuse.RenameRequest, *fuse.SetattrRequest,
				*fuse.SetxattrRequest, *fuse.StatfsRequest, *fuse.SymlinkRequest,
				*fuse.UnlockRequest:
				r.RespondError(fmt.Errorf("oops, not yet supported"))
			case *fuse.UnrecognizedRequest:
				r.RespondError(fmt.Errorf("invalid"))
			case *fuse.WriteRequest:
				res := &fuse.WriteResponse{}
				res.Size = 5
				r.Respond(res)
			}
		}
	}()

	for {
		req, err := c.ReadRequest()
		if err != nil {
			fmt.Println("error is -> ", err)
			return
		}
		if req != nil {
			reqChan <- req
		}
	}

}
