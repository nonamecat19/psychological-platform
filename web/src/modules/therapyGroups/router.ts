export const therapyGroupRouter = {
  path: '/therapy-groups',
  component: () => import('./views/TherapyGroupsLayout.vue'),
  children: [
    {
      path: ':id',
      component: () => import('./views/TherapyGroupIdPage.vue'),
    },
  ],
}
