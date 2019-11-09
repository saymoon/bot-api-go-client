package bot

import (
    "context"
    "encoding/json"
    "fmt"
)

type Snapshot struct {
    SnapshotId     string `json:"snapshot_id"`
    OpponentId     string `json:"opponent_id"`
    AssetId        string `json:"asset_id"`
    Amount         string `json:"amount"`
    TraceId        string `json:"trace_id"`
    Memo           string `json:"memo"`
    CreatedAt      string `json:"created_at"`
    CounterUserId  string `json:"counter_user_id"`
}

func SnapshotList(ctx context.Context, accessToken string, limit int, offset string, asset string)([]Snapshot, error) {
    path := fmt.Sprintf("/snapshots?limit=%d&offset=%s&asset=%s", limit, offset, asset)
    body, err := Request(ctx, "GET", path, nil, accessToken)
    if err != nil {
        return nil, err
    }
    var resp struct {
        Data  []Snapshot `json:"data"`
        Error Error      `json:"error"`
    }
    err = json.Unmarshal(body, &resp)
    if err != nil {
        return nil, BadDataError(ctx)
    }
    if resp.Error.Code > 0 {
        if resp.Error.Code == 401 {
            return nil, AuthorizationError(ctx)
        } else if resp.Error.Code == 403 {
            return nil, ForbiddenError(ctx)
        }
        return nil, ServerError(ctx, resp.Error)
    }
    return resp.Data, nil
}
