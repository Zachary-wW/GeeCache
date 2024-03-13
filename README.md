# GeeCache
GeeCache basically imitates the implementation of groupcache, and in order to limit the amount of code to about 500 lines (groupcache is about 3000 lines), some of the functions have been cut. But the overall implementation is still very close to groupcache. Supported features are:
- standalone caching and HTTP-based distributed caching
- Least Recently Used (LRU) caching policy
- Go locking mechanism to prevent cache blowouts
- Load balancing with consistent hash selection of nodes
- Optimizing inter-node binary communication using protobuf
