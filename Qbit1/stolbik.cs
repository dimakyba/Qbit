using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      string a = Console.ReadLine();
      string b = Console.ReadLine();

      while (a.Length != 10)
      {
        a = a.Insert(0, " ");
      }

      while (b.Length != 10)
      {
        b = b.Insert(0, " ");
      }

      Console.WriteLine(a);
      Console.WriteLine(b);
      Console.WriteLine("----------");

      uint numA = Convert.ToUInt32(a);
      uint numB = Convert.ToUInt32(b);
      uint result = numA + numB;

      string resultStr = result.ToString();
      while (resultStr.Length != 10)
      {
        resultStr = resultStr.Insert(0, " ");
      }
      Console.WriteLine(resultStr);
    }
  }
}
