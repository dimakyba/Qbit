using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      float beta = Convert.ToSingle(Console.ReadLine());
      int h = (int)Math.Floor(12.0f / (360.0f / beta));
      int m = (int)((12 / (360 / beta) - h) * 60);
      Console.WriteLine($"{h} {m}");




    }
  }
}
