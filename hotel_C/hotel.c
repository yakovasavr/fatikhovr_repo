/* Даны даты заезда и отъезда каждого гостя. Для каждого гостя дата заезда строго раньше даты отъезда (то есть каждый гость останавливается хотя бы на одну ночь). В пределах одного дня считается, что сначала старые гости выезжают, а затем въезжают новые. Найти максимальное число постояльцев, которые одновременно проживали в гостинице (считаем, что измерение количества постояльцев происходит в конце дня).
sample = [ (1, 2), (1, 3), (2, 4), (2, 3), ]
*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct s_list
{
                int             date;
                int             status;
                struct s_list   *next;
}               t_list;

// 0 - in, 1 - out

t_list *create_elem(int date, int status)
{
        t_list *temp;

        temp = (t_list *)malloc(sizeof(t_list));
        temp->date = date;
        temp->status = status;
        temp->next = 0;
//      printf("%d %d\n", date, status);
        return (temp);
}

void add_element(int date, int status, t_list **hotel)
{
        t_list  *temp;
        t_list  *current;
        t_list  *nextl;
        int counter;

        temp = *hotel;
        counter = 0;
        if (*hotel == 0)
                *hotel = create_elem(date, status);
        else
        {
                while (date > temp->date && temp->next != 0)
                {
                        temp = temp->next;
                        counter++;
                }
                if (temp->next == 0)
                        temp->next = create_elem(date, status);
                else
                {
                        if (counter == 0)
                        {
                                current = create_elem(date, status);
                                current->next = temp;
                                *hotel = current;

                        }
                        else
                        {
                                current = temp;
                                nextl = temp->next;
                                temp = create_elem(date, status);
                                temp->next = nextl;
                                current->next = temp;
                        }
                }
        }
}

void print_struct(t_list **hotel)
{
        t_list *temp;

        temp = *hotel;
        while (temp != 0)
        {
                printf("date is %d, status is %d\n", temp->date, temp->status);
                temp = temp->next;
        }
}

int findmaxday(t_list **hotel)
{
        int     max;
        t_list  *temp;
        int     currentdate;
        int     currentmax;
        int     maxdate;

        max = 0;
        currentmax = 0;
        temp = *hotel;
        maxdate = 0;
        while (temp->next != 0)
        {
                currentdate = temp->date;
                if (temp->status == 0)
                        currentmax++;
                else
                        currentmax--;
                if (currentmax > max && (temp->next == 0 || \
                        temp->next->date != currentdate))
                {
                        max = currentmax;
                        maxdate = temp->date;
                }
                temp = temp->next;
        }
        return (maxdate);
}

int main(int argc, char **argv)
{
        int     i;
        t_list  **hotel;
        (void)argc;
        int maxday;

        *hotel = 0;
        i = 1;
        while (argv[i])
        {
                if (i % 2 != 0)
                        add_element(atoi(argv[i]), 0, hotel);
                else
                        add_element(atoi(argv[i]), 1, hotel);
                i++;
        }
//      print_struct(hotel);
        maxday = findmaxday(hotel);
        printf("%d\n", maxday);
        return (0);
}
