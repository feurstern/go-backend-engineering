<script setup lang="ts">
import useErpLayout from './ErpLayout';

const { toggleRightDrawer, leftDrawerOpen, rightDrawerOpen } = useErpLayout();
import { Component, defineAsyncComponent } from 'vue';

const leftMenuDrawer: Component = defineAsyncComponent(
  () => import('components/erp/drawer/LeftMenuDrawer.vue'),
);

const ErpPage: Component = defineAsyncComponent(() => import('pages/ErpPage.vue'));
</script>

<template>
  <q-layout view="hHh Lpr fFf">
    <q-header elevated class="bg-green-9 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-img
            src="https://dnet.net.id/wp-content/themes/dnetv3/public/images/logo.f15d9c.png"
            style="width: 60px; height: 20px"
            fit="fill"
          />

          DAF D-net
        </q-toolbar-title>

        <q-btn dense flat round icon="menu" @click="toggleRightDrawer" />
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
      <left-menu-drawer />
    </q-drawer>

    <erp-page />
    <q-drawer show-if-above v-model="rightDrawerOpen" side="right" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <erp-page menu="finance" />
      <router-view />
    </q-page-container>

    <q-footer elevated class="bg-grey-8 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          <div>Title xxixi</div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>
