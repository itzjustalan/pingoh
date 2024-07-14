import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/tasks/new')({
  component: () => <CreateTaskPage />
})

const CreateTaskPage = () => {
  return (
  <>
    new task</>
  )
}
