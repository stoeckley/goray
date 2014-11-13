/*
** calc_cyl.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 17:51:58 2012 julien fortin
** Last update Mon Mar 12 10:18:07 2012 julien fortin
*/

#include	<math.h>
#include	<libmy.h>
#include	<rtv1.h>

t_inter *add_inter_cyl(t_scene *scene, double lambda)
{
  t_inter	*inter;

  inter = xmalloc(sizeof(*inter));
  inter->x = (lambda * scene->vx) + scene->x_oeil;
  inter->y = (lambda * scene->vy) + scene->y_oeil;
  inter->z = 0;
  return (inter);
}

double	calc_cyl(t_scene *scene, t_obj *obj)
{
  double	a;
  double	b;
  double	c;
  double	delta;
  double	lambda;

  scene = save_scene(scene, obj, 0);
  a = SQR(scene->vx) + SQR(scene->vy);
  b = 2 * ((scene->x_oeil * scene->vx) + (scene->y_oeil * scene->vy));
  c = SQR(scene->x_oeil) + SQR(scene->y_oeil) - SQR(obj->r);
  delta = SQR(b) - (4 * (a * c));
  if (delta < 0)
    {
      scene = save_scene(scene, obj, 1);
      return (-1.0);
    }
  else
    {
      lambda = (-b - sqrt(delta)) / (2 * a);
      obj->inter = add_inter_cyl(scene, lambda);
      scene = save_scene(scene, obj, 1);
      return (lambda);
    }
}
