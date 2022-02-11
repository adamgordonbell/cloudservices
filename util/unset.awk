#!/usr/bin/env -S gawk -f

# unset functions
$1 ~/\(\)/{ 
        f=substr($1, 0, index($1,"(")-1)
        s=("unset -f " f)
        print s
}
