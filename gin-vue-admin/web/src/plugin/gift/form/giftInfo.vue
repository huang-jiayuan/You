
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="礼物唯一ID:" prop="giftId">
          <el-input v-model.number="formData.giftId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="礼物名称:" prop="giftName">
          <el-input v-model="formData.giftName" :clearable="true"  placeholder="请输入礼物名称" />
       </el-form-item>
        <el-form-item label="礼物类型(1-普通,2-特效,3-稀有,4-自定义):" prop="giftType">
           <el-select v-model="formData.giftType" placeholder="请选择礼物类型(1-普通,2-特效,3-稀有,4-自定义)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in gifttypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="虚拟币价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="钻石价格(可为空):" prop="diamondPrice">
          <el-input v-model.number="formData.diamondPrice" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="礼物图标URL:" prop="iconUrl">
          <el-input v-model="formData.iconUrl" :clearable="true"  placeholder="请输入礼物图标URL" />
       </el-form-item>
        <el-form-item label="动画URL(特效礼物):" prop="animationUrl">
          <el-input v-model="formData.animationUrl" :clearable="true"  placeholder="请输入动画URL(特效礼物)" />
       </el-form-item>
        <el-form-item label="排序权重:" prop="weight">
          <el-input v-model.number="formData.weight" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否热门(1-是,0-否):" prop="isHot">
           <el-select v-model="formData.isHot" placeholder="请选择是否热门(1-是,0-否)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ishotOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="是否限时(1-是,0-否):" prop="isLimit">
           <el-select v-model="formData.isLimit" placeholder="请选择是否限时(1-是,0-否)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in islimitOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="状态(0-下架,1-上架,2-待审核):" prop="status">
           <el-select v-model="formData.status" placeholder="请选择状态(0-下架,1-上架,2-待审核)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in giftstatusOptions" :key="key" :label="item.label" :value="item.value" />
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
  createGiftInfo,
  updateGiftInfo,
  findGiftInfo
} from '@/plugin/gift/api/giftInfo'

defineOptions({
    name: 'GiftInfoForm'
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
const gifttypeOptions = ref([])
const ishotOptions = ref([])
const islimitOptions = ref([])
const giftstatusOptions = ref([])
const formData = ref({
            giftId: undefined,
            giftName: '',
            giftType: '',
            price: 0,
            diamondPrice: undefined,
            iconUrl: '',
            animationUrl: '',
            weight: undefined,
            isHot: '',
            isLimit: '',
            status: '',
        })
// 验证规则
const rule = reactive({
               giftName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               giftType : [{
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
      const res = await findGiftInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    gifttypeOptions.value = await getDictFunc('gifttype')
    ishotOptions.value = await getDictFunc('ishot')
    islimitOptions.value = await getDictFunc('islimit')
    giftstatusOptions.value = await getDictFunc('giftstatus')
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
               res = await createGiftInfo(formData.value)
               break
             case 'update':
               res = await updateGiftInfo(formData.value)
               break
             default:
               res = await createGiftInfo(formData.value)
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
