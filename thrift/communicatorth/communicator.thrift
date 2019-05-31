namespace go communicator

struct Product{
    1: double price,
    2: i32 amount,
}


service AvgService {
    double GetAvg(1: list<Product> products),
}