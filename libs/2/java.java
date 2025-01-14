public class Nonsense {
    public static void main(String[] args) {
        int a = 5;
        int b = 10;
        int c = a + b;
        System.out.println("Сумма a и b: " + c);

        for (int i = 0; i < 100; i++) {
            System.out.println("Итерация: " + i);
            if (i % 2 == 0) {
                System.out.println("Четное число: " + i);
            } else {
                System.out.println("Нечетное число: " + i);
            }
        }

        String nonsenseString = "Это бессмысленный код!";
        for (char ch : nonsenseString.toCharArray()) {
            System.out.print(ch + " ");
        }
        System.out.println();

        // Бессмысленный массив и его заполнение
        int[] nonsenseArray = new int[50];
        for (int i = 0; i < nonsenseArray.length; i++) {
            nonsenseArray[i] = (int) (Math.random() * 100);
            System.out.println("Элемент массива " + i + ": " + nonsenseArray[i]);
        }

        // Бессмысленный метод
        printNonsenseMessage();

        // Бессмысленный цикл
        for (int i = 0; i < 10; i++) {
            System.out.println("Бессмысленный цикл: " + i);
            if (i == 5) {
                System.out.println("Половина пройдена!");
            }
        }
    }

    public static void printNonsenseMessage() {
        System.out.println("Это просто бессмысленное сообщение!");
    }
}