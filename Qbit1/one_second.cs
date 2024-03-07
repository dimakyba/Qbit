using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      string[] data = Console.ReadLine().Trim().Split();
      byte h = Convert.ToByte(data[0]);
      byte m = Convert.ToByte(data[1]);
      byte s = Convert.ToByte(data[2]);
      s++;
      if (s == 60)
      {
        s = 0;
        m++;
      }
      if (m == 60)
      {
        m = 0;
        h++;
      }
      if (h == 24)
      {
        h = 0;
      }

      Console.WriteLine($"{h} {m} {s}");
    }
  }
}
