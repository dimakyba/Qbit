using System;

namespace Program
{
  class Program
  {
    static void Main(string[] args)
    {
      int n = int.Parse(Console.ReadLine());
      string[] arrString = Console.ReadLine().Trim().Split();
      int[] arr = new int[n];


      if (!isSorted(arr))
      {
        for (int i = 1; i < arr.Length; i++)
        {
          int j = i - 1;
          int key = arr[i];
          for (; j >= 0 && arr[j] > key;)
          {
            arr[j + 1] = arr[j];
            j--;
          }
          arr[j + 1] = key;
          if (arr[j + 1] != arr[i])
            Console.WriteLine(string.Join(" ", arr));
        }
      }
      else Console.WriteLine("The array is already sorted");

    }

    static bool isSorted(int[] arr)
    {
      for (int i = 0; i < arr.Length - 1; i++)
      {
        if (arr[i] > arr[i + 1]) return false;

      }
      return true;
    }
  }
}
