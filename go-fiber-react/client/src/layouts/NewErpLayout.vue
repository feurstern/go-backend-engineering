<script setup lang="ts">
import { Component, defineAsyncComponent } from 'vue';
import useNewErpLayout from './newErpLayout';

const { leftDrawerOpen, rightDrawerOpen, toggleLeftDrawer, toggleRightDrawer } = useNewErpLayout();

const leftMenuList: Component = defineAsyncComponent(() => import('components/erp/LeftMenu.vue'));
</script>

<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-green-7 text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          <q-img
            src="https://dnet.net.id/wp-content/themes/dnetv3/public/images/logo.f15d9c.png"
            style="width: 60px; height: 20px"
            fit="fill"
          />
          D-NET
        </q-toolbar-title>

        <q-btn dense flat round icon="menu" @click="toggleRightDrawer" />
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
      <!-- drawer content -->
      <left-menu-list />
    </q-drawer>

    <q-drawer show-if-above v-model="rightDrawerOpen" side="right" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer elevated class="bg-grey-8 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          <div>Title</div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>
