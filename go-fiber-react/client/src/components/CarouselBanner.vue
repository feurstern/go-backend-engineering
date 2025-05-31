<script setup lang="ts">
import { onMounted, ref, Ref } from 'vue';
import { carouselData } from './constant';
import { getUserList } from 'src/service/user';

const slide: Ref<string> = ref('style');
const fetchUserList = async () => {
  try {
    const res = await getUserList();

    if (res.success) {
      console.log('res,', res.data);
    }
  } catch (error) {
    console.log(error);
  }
};
onMounted(() => {
  fetchUserList();
});
</script>

<template>
  <div class="q-pa-md">
    <div class="q-gutter-md">
      <q-carousel
        v-model="slide"
        transition-prev="jump-right"
        transition-next="jump-left"
        swipeable
        animated
        control-color="white"
        prev-icon="arrow_left"
        next-icon="arrow_right"
        navigation-icon="radio_button_unchecked"
        navigation
        padding
        arrows
        height="300px"
        class="bg-blue text-white shadow-1 rounded-borders"
      >
        <!-- the render -->

        <q-carousel-slide
          v-for="(x, i) in carouselData"
          :key="i"
          :name="x.iconName"
          class="column no-wrap flex-center"
        >
          <q-icon name="style" size="56px" />
          <div class="q-mt-md text-center">
            {{ x.text }}
          </div>
        </q-carousel-slide>

        <!-- end  of fucking render -->
      </q-carousel>
    </div>
  </div>
</template>
