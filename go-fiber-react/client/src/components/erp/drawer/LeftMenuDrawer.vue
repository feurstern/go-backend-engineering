<script setup lang="ts">
import useLeftMenuDrawer from './LeftMenuDrawer';
import { LeftMenuDrawerList } from 'src/constant';

const { menu, handleMenuChange } = useLeftMenuDrawer();
</script>

<template>
  <div class="q-pa-md">
    <q-list padding class="rounded-borders text-primary">
      <q-item
        v-for="(x, i) in LeftMenuDrawerList"
        clickable
        :key="i"
        v-ripple
        :active="menu === x.name"
        @click="handleMenuChange(x.name)"
        active-class="my-menu-link"
      >
        <q-item-section avatar>
          <q-icon :name="x.icon" />
        </q-item-section>

        <q-item-section>{{ x.name.toLocaleUpperCase() }}</q-item-section>
      </q-item>
    </q-list>
  </div>
</template>
