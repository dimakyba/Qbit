using System;

class Program
{
  static void Main()
  {

    char c = Convert.ToChar(Console.ReadLine());

    if (c >= 65 && c <= 90)
    {
      c = (char)(c + 32);
      Console.WriteLine(c);
    }
    else if (c >= 97 && c <= 122)
    {
      c = (char)(c - 32);
      Console.WriteLine(c);
    }
  }
}
