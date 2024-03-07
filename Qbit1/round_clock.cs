using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      string[] data = Console.ReadLine().Trim().Split();
      int h = Convert.ToInt32(data[0]);
      int m = Convert.ToInt32(data[1]);
      int s = Convert.ToInt32(data[2]);
      float result = (float)(30.0 * h + 0.5 * m + (1.0 / 120) * s);
      Console.WriteLine(result);
    }
  }
}
