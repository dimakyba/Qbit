using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int n = int.Parse(Console.ReadLine());
        int[] array = Console.ReadLine().Split(' ').Select(int.Parse).ToArray();

        int L = GetFirstNegativeIndex(array);
        int R = GetLastNegativeIndex(array);

        if (L != -1 && R != -1 && L < R)
        {
            int sum = CalculateSumBetweenIndices(array, L, R);
            Console.WriteLine(sum);
        }
        else
        {
            int sum = SumArray(array);
            Console.WriteLine(sum);
        }
    }

    static int GetFirstNegativeIndex(int[] arr)
    {
        for (int i = 0; i < arr.Length; i++)
        {
            if (arr[i] < 0)
            {
                return i;
            }
        }

        return -1;
    }

    static int GetLastNegativeIndex(int[] arr)
    {
        for (int i = arr.Length - 1; i >= 0; i--)
        {
            if (arr[i] < 0)
            {
                return i;
            }
        }

        return -1;
    }

    static int CalculateSumBetweenIndices(int[] arr, int L, int R)
    {
        int sum = 0;

        if (L > R)
        {
            (L,R) = (R,L);
        }

        for (int i = L; i <= R; i++)
        {
            sum += arr[i];
        }

        return sum;
    }

    static int SumArray(int[] arr)
    {
        int sum = 0;
        foreach (int num in arr)
        {
            sum += num;
        }
        return sum;
    }
}
