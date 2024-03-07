using System;

class Program
{
  static void Main(string[] args)
  {
    string p = Console.ReadLine();
    string t = Console.ReadLine();
    int n = int.Parse(Console.ReadLine()) - 1;
    int pLen = p.Length;
    int tLen = t.Length;
    int iterations = 0; // Переменная для подсчета итераций

    for (int i = 0; i <= tLen - pLen; i++)
    {
      iterations++; // Увеличиваем счетчик итераций при каждой итерации цикла
      string s = t.Substring(n + i, pLen);
      if (s == p)
      {
        Console.WriteLine("YES");
        Console.WriteLine(iterations); // Выводим количество итераций
        return; // Выходим из программы после нахождения подстроки
      }
    }

    Console.WriteLine("NO");
    Console.WriteLine(iterations); // Выводим количество итераций, если подстрока не найдена
  }
}
