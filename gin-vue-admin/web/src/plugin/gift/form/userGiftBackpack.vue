
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="所属用户ID，关联 user.id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="礼物ID:" prop="giftId">
          <el-input v-model.number="formData.giftId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="持有数量:" prop="count">
          <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="获得来源:" prop="obtainSource">
           <el-select v-model="formData.obtainSource" placeholder="请选择获得来源" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in obtainsourceOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="过期时间(NULL-永久):" prop="expireTime">
          <el-date-picker v-model="formData.expireTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="是否有效(1-是,0-否):" prop="isValid">
           <el-select v-model="formData.isValid" placeholder="请选择是否有效(1-是,0-否)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in is_validOptions" :key="key" :label="item.label" :value="item.value" />
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
  createUserGiftBackpack,
  updateUserGiftBackpack,
  findUserGiftBackpack
} from '@/plugin/gift/api/userGiftBackpack'

defineOptions({
    name: 'UserGiftBackpackForm'
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
const obtainsourceOptions = ref([])
const is_validOptions = ref([])
const formData = ref({
            userId: undefined,
            giftId: undefined,
            count: undefined,
            obtainSource: '',
            expireTime: new Date(),
            isValid: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               giftId : [{
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
      const res = await findUserGiftBackpack({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    obtainsourceOptions.value = await getDictFunc('obtainsource')
    is_validOptions.value = await getDictFunc('is_valid')
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
               res = await createUserGiftBackpack(formData.value)
               break
             case 'update':
               res = await updateUserGiftBackpack(formData.value)
               break
             default:
               res = await createUserGiftBackpack(formData.value)
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
