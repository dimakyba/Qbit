using System;
using System.Text;

class Program
{
  static void Main(string[] args)
  {
    string s = Console.ReadLine();
    uint counter = 0;
    char[] chars = { '.', ',', '!', ';', ':' };
    foreach (char c in s)
    {
      if (Array.IndexOf(chars, c) != -1)
      {
        counter++;
      }

    }
    System.Console.WriteLine(counter);
  }
}
