import type { TodoApi } from './todoApi'
import type { Todo } from '@/types'

let todos: Todo[] = [
  { id: 1, text: "コーヒー豆を買う", done: false },
  { id: 2, text: "日記を書く", done: false },
  { id: 3, text: "散歩する", done: true },
]
let nextId = 4

export const todoApiMock: TodoApi = {
    getTodos: async () => [...todos],

    addTodo: async (text) => {
        const todo = { id: nextId++, text: text, done: false }
        todos.push(todo)
        return todo
    },

    updateTodo: async (id, patch) => {
        const todo = todos.find(t => t.id === id)
        if (!todo) throw new Error('Not found')
        if (patch.text !== undefined) todo.text = patch.text
        if (patch.done !== undefined) todo.done = patch.done
        return { ...todo }
    },

    deleteTodo: async (id) => {
        todos = todos.filter(t => t.id !== id)
    },
}
