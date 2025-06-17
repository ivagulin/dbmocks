for i in `seq -w 1 100`; do sed "s/Suite1/Suite$i/g" template.go > suite"$i"_test.go ; done
