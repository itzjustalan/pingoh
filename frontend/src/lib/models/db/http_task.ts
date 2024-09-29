import { z } from "zod";

const zodStringOneOf = (values: string[]) =>
    z.string().refine((v) => values.includes(v));

export const httpTaskSchema = z
    .object({
        "accepted_status_codes": z.string(),
        "auth_method": zodStringOneOf(['none',
            'basic',
            'oauth2']),
        "body": z.string(),
        "encoding": zodStringOneOf(['none',
            'text',
            'html',
            'form',
            'json',
            'xml']),
        "headers": z.record(z.string(), z.string()),
        "id": z.number(),
        "method": zodStringOneOf(['GET',
            'POST',
            'PUT',
            'PATCH',
            'DELETE',
            'HEAD',
            'OPTIONS']),
        "retries": z.number(),
        "task_id": z.number(),
        "timeout": z.number(),
        "url": z.string().url(),
    })
    .strict();
export type HttpTask = z.infer<typeof httpTaskSchema>;