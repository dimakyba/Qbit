using System;

class Program
{
    static void Main()
    {
        int N = int.Parse(Console.ReadLine());

        int[] array = new int[N];
        string[] elements = Console.ReadLine().Split();

        for (int i = 0; i < N; i++)
        {
            array[i] = int.Parse(elements[i]);
        }

        FindSecondLargest(array, out int secondLargestElement, out int secondLargestIndex);

        Console.WriteLine($"{secondLargestElement} {secondLargestIndex}");
    }

    static void FindSecondLargest(int[] arr, out int secondLargestElement, out int secondLargestIndex)
    {
        int maxElement = int.MinValue;
        int maxIndex = -1;

        for (int i = 0; i < arr.Length; i++)
        {
            if (arr[i] > maxElement)
            {
                maxElement = arr[i];
                maxIndex = i;
            }
        }

        arr[maxIndex] = int.MinValue;

        secondLargestElement = int.MinValue;
        secondLargestIndex = -1;

        for (int i = 0; i < arr.Length; i++)
        {
            if (arr[i] > secondLargestElement)
            {
                secondLargestElement = arr[i];
                secondLargestIndex = i;
            }
        }

        secondLargestIndex++;
    }
}
