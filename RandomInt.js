# フィッシャー–イェーツ

# 配列の要素はいくつか
array_num = 3;

let hosts_num_array = Array.from(new Array(array_num)).map((v,i) => i)
let new_array = [];
let i = 1;

while (i < 10) {
  console.log(i + "回目")
  randomInt(hosts_num_array, new_array)
  if (hosts_num_array.length === 0) {
    Array.prototype.push.apply(hosts_num_array, new_array);
    new_array = [];
  }
  i++
}

function randomInt(hosts_num_array, new_array) {
    n = hosts_num_array.length;
    k = Math.floor(Math.random() * n);
    j = hosts_num_array[k];

    new_array.push(hosts_num_array[k]);
    hosts_num_array.splice(k, 1);
    console.log(hosts_num_array);
    console.log(new_array);
    // return j
}
