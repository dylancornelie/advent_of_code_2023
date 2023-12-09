#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <stdint.h>
#include <ctype.h>
// For input file
#define MAX_LINE_LENGTH 141
#define NUMBER_OF_LINE 139
// For sample file
// #define MAX_LINE_LENGTH 11
// #define NUMBER_OF_LINE 9

int find_min_in_array(int *array, int size);
int find_max_in_array(int *array, int size);
void reset_array(int *array, int size);
bool has_adjacent_symbol(int pos_i, int *pos_j, int pos_j_size, char **data);
int get_sum_of_adjacent_numbers(int pos_i, int pos_j, char **data);


int main()
{
    const int PUZZLE_PART = 2;
    FILE *file = fopen("./input.txt", "r");
    // FILE *file = fopen("./sample.txt", "r");

    char line[MAX_LINE_LENGTH];
    char **data = calloc(NUMBER_OF_LINE, MAX_LINE_LENGTH * sizeof(char));
    assert(data != NULL);
    int current_line = 0;

    // Loading each line of the file into memory
    while (fscanf(file, "%[^\n]\n", line) != -1)
    {
        data[current_line] = calloc(1, MAX_LINE_LENGTH * sizeof(char));
        assert(data[current_line] != NULL);
        strncpy(data[current_line], line, MAX_LINE_LENGTH);
        current_line++;
    }

    int result = 0;
    int indexes_of_number[MAX_LINE_LENGTH] = {};
    reset_array(indexes_of_number, MAX_LINE_LENGTH);
    bool found_a_number = false;
    int current_index = 0;
    for (int i = 0; i <= NUMBER_OF_LINE; i++)
    {
        for (int j = 0; j < MAX_LINE_LENGTH; j++)
        {
            if (PUZZLE_PART == 1)
            {
                if (isnumber(data[i][j]))
                {
                    found_a_number = true;
                    // Save all the indexes of each character composing the number
                    indexes_of_number[current_index++] = j;
                }
                else if (found_a_number)
                {
                    if (has_adjacent_symbol(i, indexes_of_number, current_index, data))
                    {
                        // Create a char buffer containing the number
                        char number_as_string[current_index + 1];
                        for (int y = 0; y < current_index; y++)
                        {
                            number_as_string[y] = data[i][indexes_of_number[y]];
                        }
                        // Converting the number into an int
                        result += atoi(number_as_string);
                    }
                    // Reset the states
                    found_a_number = false;
                    current_index = 0;
                    reset_array(indexes_of_number, MAX_LINE_LENGTH);
                }
            }
            else if ((int)data[i][j] == 42)
            {
                result += get_sum_of_adjacent_numbers(i, j, data);
            }
        }
    }

    // Memory clean up
    for (int i = 0; i < NUMBER_OF_LINE; i++)
    {
        free(data[i]);
    }
    free(data);
    fclose(file);
    printf("Result is %d\n", result);
    return 0;
}

bool has_adjacent_symbol(int pos_i, int *pos_j, int pos_j_size, char **data)
{
    int i = pos_i - 1 < 0 ? pos_i : pos_i - 1;
    int max_i = pos_i + 1 > NUMBER_OF_LINE ? pos_i : pos_i + 1;
    int min_pos_j = find_min_in_array(pos_j, pos_j_size);
    int max_pos_j = find_max_in_array(pos_j, pos_j_size);
    int min_j = min_pos_j - 1 < 0 ? min_pos_j : min_pos_j - 1;
    int max_j = max_pos_j + 1 > MAX_LINE_LENGTH ? max_pos_j : max_pos_j + 1;
    for (; i <= max_i; i++)
    {
        for (int j = min_j; j <= max_j; j++)
        {
            int ascii_value = data[i][j];
            if (!isnumber(data[i][j]) && ascii_value != 46 && ascii_value != 0)
            {
                return true;
            }
        }
    }
    return false;
}

int get_sum_of_adjacent_numbers(int pos_i, int pos_j, char **data)
{
    int i = pos_i - 1 < 0 ? pos_i : pos_i - 1;
    int max_i = pos_i + 1 > NUMBER_OF_LINE ? pos_i : pos_i + 1;
    int min_j = pos_j - 1 < 0 ? pos_j : pos_j - 1;
    int max_j = pos_j + 1 > MAX_LINE_LENGTH ? pos_j : pos_j + 1;
    int gear_ratio = 0;
    bool is_a_valid_gear = false;

    for (; i <= max_i; i++)
    {
        int upper_index = 1;
        for (int j = min_j; j <= max_j; j++)
        {
            if (isnumber(data[i][j]) && j > upper_index)
            {
                int lower_index = j;
                while (lower_index >= 0 && isnumber(data[i][lower_index]))
                {
                    --lower_index;
                }
                lower_index++;
                upper_index = j;
                while (lower_index <= MAX_LINE_LENGTH && isnumber(data[i][upper_index]))
                {
                    ++upper_index;
                }
                char number_as_string[4] = {};
                for (int y = 0; y < upper_index - lower_index; y++)
                {
                    number_as_string[y] = data[i][lower_index + y];
                }
                if (gear_ratio == 0)
                {
                    gear_ratio = atoi(number_as_string);
                }
                else
                {
                    is_a_valid_gear = true;
                    gear_ratio *= atoi(number_as_string);
                }
            }
        }
    }
    return is_a_valid_gear ? gear_ratio : 0;
}

int find_min_in_array(int *array, int size)
{
    int min = INT16_MAX;
    for (int i = 0; i < size; i++)
    {
        if (array[i] < min)
        {
            min = array[i];
        }
    }
    return min;
}

int find_max_in_array(int *array, int size)
{
    int max = INT16_MIN;
    for (int i = 0; i < size; i++)
    {
        if (array[i] > max)
        {
            max = array[i];
        }
    }
    return max;
}

void reset_array(int *array, int size)
{
    for (int i = 0; i < size; i++)
    {
        array[i] = 0;
    }
    return;
}