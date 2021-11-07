// Normal code for computing Inverse product of a matrix.
func InverseProduct(a Matrix, b Matrix) {
  a_inv := Inverse(a)
  b_inv := Inverse(b)
  return Product(a_inv, b_inv)
}


// With Futures... 
func InverseProduct(a Matrix, b Matrix) {
  a_inv_future := InverseFuture(a) // started as a goroutine
  b_inv_future := InverseFuture(b) // started as a goroutine
  a_inv := <-a_inv_future
  b_inv := <-b_inv_future
  return Product(a_inv, b_inv)
}

func InverseFuture(a Matrix) (chan Matrix) {
  future := make(chan Matrix)
  go func() { future <- Inverse(a) }()
  return future
}
