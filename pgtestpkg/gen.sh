for i in `seq -w 1 100`; do sed "s/Suite1/Suite$i/g" template.gotmpl > suite"$i"_test.go ; done
