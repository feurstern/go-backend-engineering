<script lang="ts" setup>
import useLoginComponent from './LoginComponent';
const { email, password, onReset, handleLogin } = useLoginComponent();
</script>

<template>
  <div class="q-pa-md" style="max-width: 450px">
    <q-form class="q-gutter-md" @submit="handleLogin" @reset="onReset">
      <q-input
        filled
        label="email"
        v-model="email"
        hint="naziincolour123@gmail.com"
        lazy-rules
        :rules="[(val) => (val && val.length > 0) || 'Please input your email correctly!']"
      />

      <q-input
        filled
        label="password"
        type="password"
        v-model="password"
        hint="password"
        :rules="[(val) => (val && val.length > 3) || 'Please input the password correctly']"
      />

      <div>
        <q-btn label="Submit" type="submit" color="primary" />
        <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
      </div>
    </q-form>
  </div>
</template>
