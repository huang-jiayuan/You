
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标签名称:" prop="tagName">
          <el-input v-model="formData.tagName" :clearable="true"  placeholder="请输入标签名称" />
       </el-form-item>
        <el-form-item label="标签描述:" prop="tagDesc">
          <el-input v-model="formData.tagDesc" :clearable="true"  placeholder="请输入标签描述" />
       </el-form-item>
        <el-form-item label="状态(0-启用,1-停用):" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态(0-启用,1-停用)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createRoomTagDict,
  updateRoomTagDict,
  findRoomTagDict
} from '@/plugin/chamber/api/roomTagDict'

defineOptions({
    name: 'RoomTagDictForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const stateOptions = ref([])
const formData = ref({
            tagName: '',
            tagDesc: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               tagName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tagDesc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRoomTagDict({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stateOptions.value = await getDictFunc('state')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createRoomTagDict(formData.value)
               break
             case 'update':
               res = await updateRoomTagDict(formData.value)
               break
             default:
               res = await createRoomTagDict(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
