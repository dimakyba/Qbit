using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      uint n = UInt32.Parse(Console.ReadLine());
      int i = 0;
      Console.Write(n * 4990);
      double k = (double)(n * (249.5 / 3e-23));
      while (k > 10)
      {
        k /= 10;
        i++;
      }
      k = Math.Round(k, 3);
      Console.WriteLine($" {k}E+{i}");
    }
  }
}
