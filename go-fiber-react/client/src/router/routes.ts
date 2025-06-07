import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: '/test', component: () => import('pages/TestingPage.vue') },
      { path: '/registration', component: () => import('pages/RegistrationPage.vue') },
    ],
  },

  {
    path: '/erp',
    component: () => import('layouts/NewErpLayout.vue'),
    children: [
      { path: 'registration', component: () => import('pages/RegistrationPage.vue') },
      {
        path: 'finance',
        component: () => import('pages/ERP/FinancePage.vue'),
      },
      {
        path: 'admin',
        component: () => import('pages/ERP/AdminPage.vue'),
      },
      {
        path: 'marketing',
        component: () => import('pages/ERP/MarketingPage.vue'),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
