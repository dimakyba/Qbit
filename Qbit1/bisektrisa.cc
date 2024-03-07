#include <iostream>
#include <cmath>




int main()
{
    double a, b,c;

    std::cin >> a;
    std::cin >> b;
    std::cin >> c;

    bool isTrue = 0 < a < 10000 && 0 < c < 10000 && 0 < b < 10000 && a + b > c && a + c > b && b + c > a;

    double p = (a + b + c) / 2;
    if (isTrue)
    {
        double lA = ( 2 * sqrt(b * c * p *(p - a) ) / (b + c));
        double lB = ( 2 * sqrt(a * c * p *(p - b) ) / (a + c));
        double lC = ( 2 * sqrt(a * b * p *(p - c) ) / (b + a));

        std::cout << lA << std::endl;
        std::cout << lB << std::endl;
        std::cout << lC << std::endl;


    }





    return 0;

}
