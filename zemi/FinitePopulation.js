// 20210708にゼミでやった。
// 母数に対して世論調査でどのくらいの人数調査すべきか出す。

parameter = 16000
f_child = 0.05 / 1.96
s_child = (parameter - 1) / 0.25

children = f_child * f_child * s_child + 1

console.log(parameter / children)