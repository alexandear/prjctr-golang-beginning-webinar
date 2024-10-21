## In C-files dir
```shell
gcc -fPIC -shared -o libmathlib.so mathlib.c
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:.
```

