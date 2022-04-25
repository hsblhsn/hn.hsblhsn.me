import { CommentConnection } from './Comment'
import { OpenGraph } from './OpenGraph'
import { User } from './User'

export interface Story {
  by: User
  descendants: number
  id: string
  kids: number[]
  score: number
  time: number
  title: string
  text: string
  type: string
  url: string
  comments?: CommentConnection
  openGraph?: OpenGraph
  html?: string
}
