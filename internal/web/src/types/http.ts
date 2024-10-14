/**
 * Res response
 */
export type Res<T> = {
    success: boolean;
    message: string;
    data: T;
};

/**
 * Pagination is a metadata for pagination
 */
export type Pagination = {
    total_records: number;
    limit_records: number;
    current_page: number;
    total_pages: number;
    next_page: number | null;
    prev_page: number | null;
};

/**
 * ResWithPagination response with pagination
 */
export type ResWithPagination<T> = {
    success: boolean;
    message: string;
    data: T[];
    pagination: Pagination;
};
