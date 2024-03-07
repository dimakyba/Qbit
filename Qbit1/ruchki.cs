using System;

class Program
{
  static void Main()
  {
    string[] data = Console.ReadLine().Trim().Split();
    int a = Int32.Parse(data[0]);
    int b = Int32.Parse(data[1]);
    string[] data2 = Console.ReadLine().Trim().Split();
    int c = Int32.Parse(data2[0]);
    int d = Int32.Parse(data2[1]);
    int pens = 0;


    int priceInKops = a * 100 + b;
    int cashInKops = c * 100 + d;

    while (cashInKops >= priceInKops)
    {
      cashInKops -= priceInKops;
      pens++;
    }
    Console.WriteLine(pens);
    Console.WriteLine($"{cashInKops / 100} {cashInKops % 100}");
  }
}
