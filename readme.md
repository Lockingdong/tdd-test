# 等差級數

## 算出項數為 5 的等差級數的總和
- 輸入`首項`和`末項`，算出項數為5的等差級數的總和
- 當`首項`和`末項`為0時，必須回傳錯誤訊息 `firstTerm and lastTerm cannot be zero`

## 算出項數為 10 的等差級數的總和
- 輸入`首項`和`末項`，算出項數為10的等差級數的總和
- 當`首項`和`末項`為0時，必須回傳錯誤訊息 `firstTerm and lastTerm cannot be zero`

# Black Jack
牌點數 (無分花色)
A -> 1 or 11
2 -> 2
3 -> 3
4 -> 4
5 -> 5
6 -> 6
7 -> 7
8 -> 8
9 -> 9
10 -> 10
J -> 10
Q -> 10
K -> 10

## Session 1 - 計算點數:
測項：
2, 3 = 5
2, 3, 4 = 9
A = 11
A, A = 12
A, J = 21
A, J, K = 21
A, A, A = 13
A, K, 2 = 13
J, Q, 5 = 25

hint: 
`計算` `手牌`中，所有`卡片`的點數 
