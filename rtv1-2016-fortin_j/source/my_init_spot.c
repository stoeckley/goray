/*
** my_init_spot.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sat Mar 10 14:08:04 2012 julien fortin
** Last update Mon Mar 12 10:45:35 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_spot	*my_init_spot()
{
  t_spot	*spot;

  spot = xmalloc(sizeof(*spot));
  spot->x = 3000.0;
  spot->y = 5000.0;
  spot->z = -1000.0;
  spot->color = 0x00AAAAAA;
  spot->next = NULL;
  return (spot);
}