# What
Custom clone of `wc` cli tool in golang

# How to Run
1. Clone the git repository
2. Run `go mod tidy`
3. Run `go install ccwc`
4. The binaries will be present in the GOPATH

# Commands
    ccwc [flag] [filepath]

# Flag Supported
    c : count of bytes in the file
    w : count of words in the file
    l : count of lines in the file
    m : count of multibytes in the file


