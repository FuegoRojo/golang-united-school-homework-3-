package homework

func average(input [15]float32) (result float32) {
        var sum float32 = 0
        var item_qty float32 = 0
        for _, myitem := range input{
                sum = sum+myitem
                item_qty = item_qty + 1
        }
        //print(int(item_qty))
        //print(" ")
        //Place your code here
        return sum/float32(item_qty)

}
