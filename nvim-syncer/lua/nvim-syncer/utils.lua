local M = {}

function M.ser(data)
    return vim.json.encode(data)
end

function M.deser(data)
    return vim.json.decode(data)
end


return M
