# **RADLOG** 
Radical logging framework

## Rationale 
Radlog is an experimental logging library vaguely inspired by log4j and based on few alternative principles:
- The app is responsible for both log generation and filtering
- Log generation should be strongly opinionated and as standard as possible (log all worry after)
- Log filtering should be configurable at function level
- Logging should have the smallest impact on performance
- Radlog trades performance for observarbility

## Functionality 
Radlog automatically generates a radix tree of logging nodes at startup, binding the lognode to the call stack location and caching the nodes for subsequent calls
Every log entry generated around the code is bound to a node in the radix tree
Filtering can be applied to subbranches based on rules

## **Example** 

*Call stack:* 
```
.
├── main()                                                  l.inf, l.err
    ├── u := newuser()                                             l.err 
    ├── n := newbook()                                             l.err 
    └── u.readbook(n)                                              l.err 
             ├── adduserbooks(u,n)                          l.inf 
             │        └── insertdbentry(u,n)         l.dbg, l.inf
             └── removebookfromlibrary(n)                   l.inf 
                      └── insertdbentry(u,n)         l.dbg, l.inf
```

*Log nodes:*  
```
.main (l.inf, l.err)
.main.newuser  (l.err)
.main.newbook  (l.err)
.main.readbook (l.err)
.main.readbook.adduserbooks (l.inf)
.main.readbook.adduserbooks.insertdbentry (l.dbg,l.inf)
.main.readbook.adduserbooks (l.inf)
.main.readbook.removebookfromlibrary.insertdbentry (l.dbg,l.inf)
```

*Rules configuration examples:*  
```
# debug all newuser() subfunctions
.main=inf
.main.newuser=dbg

# debug insertdbentry() inside adduserbooks() function alone*
.main=inf
.main.readbook.adduserbooks.insertdbentry=dbg

# debug adduserbooks() function with no sublogs*
.main=inf
.main.readbook.adduserbooks=dbg
.main.readbook.adduserbooks.*=inf
```

