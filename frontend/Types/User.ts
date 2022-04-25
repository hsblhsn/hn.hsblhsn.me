import { ConnectionT, Item } from './Node'

export interface User {
  id: string
  created?: number
  about?: string
  karma?: number
  submitted?: ConnectionT<Item>
}
