<script setup lang="ts">
import { LeftMenuDrawerList } from 'src/constant';
import useLeftMenu from './LeftMenu';

const { handleChange, selectedLeftMenu } = useLeftMenu();
</script>

<template>
  <div class="q-pa-md" style="max-width: 500px">
    <q-list bordered>
      <q-item
        @click="handleChange(x.name)"
        clickable
        v-ripple
        :active="selectedLeftMenu.toLocaleLowerCase() === x.name.toLocaleLowerCase()"
        v-for="(x, i) in LeftMenuDrawerList"
        :key="i"
        active-class="my-menu-link"
      >
        <q-item-section avatar>
          <q-icon :name="x.icon" color="primary" />
        </q-item-section>
        <q-item-section>
          {{ x.name }}
        </q-item-section>
      </q-item>

      <q-item clickable v-ripple>
        <q-item-section avatar>
          <q-icon color="primary" name="bluetooth" />
        </q-item-section>

        <q-item-section>Icon as avatar</q-item-section>
      </q-item>

      <q-item clickable v-ripple>
        <q-item-section avatar>
          <q-avatar color="teal" text-color="white" icon="bluetooth" />
        </q-item-section>

        <q-item-section>Avatar-type icon</q-item-section>
      </q-item>
    </q-list>
  </div>
</template>

<style scoped lang="sass">
.my-menu-link
  color: white
  background: #F2C037
</style>
