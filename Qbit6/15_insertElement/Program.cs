using System;
using System.Linq;

namespace Program
{
  class Program
  {
    static void Main()
    {
      int length = int.Parse(Console.ReadLine());
      int[] array = Console.ReadLine().Split().Select(int.Parse).ToArray();
      string[] data = Console.ReadLine().Split();
      int index = int.Parse(data[0]) - 1;
      int number = int.Parse(data[1]);
      Array.Resize(ref array, array.Length + 1);
      for (int i = array.Length - 1; i > index; i--)
      {
        array[i] = array[i - 1];
      }
      array[index] = number;
      Console.WriteLine(string.Join(" ", array));
    }
  }

}
