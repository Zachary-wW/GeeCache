package geecache

import pb "geecache/geecachepb"

// input the key and choose the PeerGetter
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// use Get method to find result in group, as the http client
type PeerGetter interface {
	// Get(group string, key string) ([]byte, error)
	Get(in *pb.Request, out *pb.Response) error
}
