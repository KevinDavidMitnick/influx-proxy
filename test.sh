#!/bin/sh
for i in `seq 1 1`;do
    echo $i
    timestamp=$(( 1530842400099987666 + $i ))
    curl -i -XPOST 'http://10.21.1.215:9086/write?db=test' --data-binary "weather,altitude=888,area='liucong' temperature=88,humidity=${i} ${timestamp}"
done

