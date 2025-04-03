export type Environment = {
    id: int
    description?: string
    name: string
    tags: string[]
    base_url: string
    created_at: Date
    updated_at?: Date

    sessions?: Session[]
    requests?: Request[]
}

export type Session = {
    id: int
    environment_id: int
    user_id: string
    created_at: Date
    updated_at?: Date
    tags: string[]

    requests?: Request[]
    environment?: Environment
}

export type Request = {
    id: int
    environment_id: int
    session_id: string
    user_id: string
    method: string
    name: string
    path: string
    headers?: string
    body?: string
    called_at: Date
    created_at: Date
    updated_at?: Date

    Session?: Session
    Environment?: Environment
}