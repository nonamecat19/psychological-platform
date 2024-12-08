export const usersRouter = {
  path: '/users',
  component: () => import('./views/UsersLayout.vue'),
  children: [
    {
      path: 'psychologists',
      component: () => import('./views/UsersPsychologistsView.vue'),
    },
  ],
}
