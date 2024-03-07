using System;
using System.Text;

class Program
{
  static void Main(string[] args)
  {
    string s = Console.ReadLine();
    StringBuilder sb = new StringBuilder(s);

    int i = 0;

    foreach (char c in sb.ToString())
    {
      if (c == 'z')
      {
        sb.Insert(i, 'z');
        i++;
      }
      i++;
    }

    Console.WriteLine(sb);
  }
}
