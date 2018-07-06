#!/bin/sh
for i in `seq 1 5000`;do
    echo $i
    timestamp=$(( 1530842400099987666 + $i ))
    curl -i -XPOST 'http://10.21.1.215:6666/write?db=testDB' --data-binary "cpu,altitude=888,area='liucong' temperature=88,humidity=${i} ${timestamp}"
done

