# assignment-go
* 「assignment」ディレクトリで「go run .」を実行

## 課題１
* /api/categories：カテゴリの配列を含むJSONを返します。これらのカテゴリにはlorem ipsumを使用するか、ランダムに文字列を生成するか、あるいは固定の設定ファイルを使用することができます。
  * 「go mod init example.com/myapp」でmodファイルを作成
  * 「go get github.com/drhodes/golorem」golorem を module に追加
  * ブラウザ等で「 http://localhost:8000/api/categories 」にアクセス

## 課題２
* /api/calculator?o={operator}&x={x}&{y}：指定された2つのオペランド（xとy）に対して、指定された演算子（o）を使って計算を行い、その結果を返します。サポートされる演算子には、加算（+）、減算（-）、乗算（*）、除算（/）があります。
  * {operator}は「add(+)」「sub(-)」「mul(*)」「div(/)」
  * アクセス例「 http://localhost:8000/api/calculator?o=add&x=1&y=2 」