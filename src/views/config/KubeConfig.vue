<script setup lang="ts">
import {reactive, ref} from 'vue'
import {updateKubeconfigContent} from "@/api/kubeconfig";
import axios from "axios";
import router from "@/router";


interface KubeConfig {
  content: string;
}

let obj: KubeConfig = {
  content: ""
}

const formModel = reactive(obj)

const onSubmit = () => {
  updateKubeconfigContent(formModel).then(res => {
    console.log(res)
    router.push({path: '/'})
  })
}

</script>

<template>
  <div>
    <el-form v-model="formModel">
      <el-form-item label="请输入kubeConfig文件内容">
        <el-input id="input"
                  v-model="formModel.content"
                  :rows="10"
                  type="textarea"
                  placeholder="Please input"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Submit</el-button>
      </el-form-item>
    </el-form>

  </div>
</template>

<style scoped lang="scss">


</style>
